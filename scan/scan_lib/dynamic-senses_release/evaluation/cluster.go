package evaluation

import (
  "fmt"
  "sort"
  "encoding/gob"
  "bytes"
  "os"
  "dynamic-senses_release/util"
  "io/ioutil"
  "strings"
)


type cluster struct {
  Name int
  Members []string
}



type NounCluster map[string]map[int]int



/* create (cluster-indexed) representation of gold cluster used by purity/collocation and v-measure */
func getGoldCluster(goldFile string) (gold map[int]map[string]float64, instances int) {
  // read data
  goldB, _ := ioutil.ReadFile(goldFile)
  goldStr := string(goldB)
  goldCL := strings.Split(goldStr, "\n\n") /*"====" for chinese*/
  // create gold struct
  gold = make(map[int]map[string]float64)
  for idx, cluster := range(goldCL) {
    words := strings.Split(cluster, "\n")[1:]
    instances += len(words)
    gold[idx] = make(map[string]float64)
    for _,word := range(words) {
      gold[idx][strings.ToLower(strings.TrimSpace(word))]=1.0
    }
  }
  return
}


func GetGoldClusterNameMap(goldFile string) (map[int]string) {
  // read data
  goldB, _ := ioutil.ReadFile(goldFile)
  goldStr := string(goldB)
  goldCL := strings.Split(goldStr, "\n\n")
  // create gold struct
  names := make(map[int]string)
  for idx, cluster := range(goldCL) {
    names[idx]=strings.Split(cluster, "\n")[0]
  }
  return names
}


/* create FUZZY (cluster-indexed) representation of gold cluster used by purity/collocation and v-measure
 * BY DISTRIBUTING MASS EQUALLY ACROSS ALL CLASSES A DATA POINT BELONGS TO */
func get_fuzzy_goldCluster(goldFile string) (gold map[int]map[string]float64, instances int) {
  // read data
  goldB, _ := ioutil.ReadFile(goldFile)
  goldStr := string(goldB)
  goldCL := strings.Split(goldStr, "\n\n") /*"====" for chinese*/
  // create word:ambiguity map
  ambiguity_map := make(map[string]float64)
  for _, cluster := range(goldCL) {
    words := strings.Split(cluster, "\n")[1:]
    for _,word := range(words) {
      ambiguity_map[strings.TrimSpace(word)]++
    }
  }
  // create gold struct
  gold = make(map[int]map[string]float64)
  for idx, cluster := range(goldCL) {
    words := strings.Split(cluster, "\n")[1:]
    instances += len(words)
    gold[idx] = make(map[string]float64)
    for _,word := range(words) {
      gold[idx][strings.ToLower(strings.TrimSpace(word))]=1.0/ambiguity_map[strings.TrimSpace(word)]
    }
  }
  return
}


/* create (cluster-indexed) representation of induced clusters used by purity/collocation and v-measure */
func getSystemCluster(data map[string]map[int]int, mode string) (topics map[int]map[string]float64) {
  topics = make(map[int]map[string]float64)
  for word, tpCounts := range(data) {
    norm := 0
    for _,cnt := range(tpCounts) {
      norm+=cnt
    }
    for topic,cnt := range(tpCounts) {
      if _,ok := topics[topic] ; !ok {
	topics[topic] = make(map[string]float64)
      }
      if mode == "norm" {
	topics[topic][word] = float64(cnt) / float64(norm)
      } else {
	topics[topic][word] = float64(cnt)
      }
    }
  }
  return
}




/* Read cluster from .txt file as created by CHINESE WHISPERS
 * and create NounCluster representation 
 * Potentially not all TARGET terms are contained in those clusters 
 * because of the edge weight parameter --> add them back as singletons! */
func ReadCluster(file string, dict map[string]int) (cluster NounCluster) {
  tmp:=make(map[string]int,len(dict))
  for k,v := range(dict) {
    tmp[k]=v
  }
  cluster = NounCluster{}
  binCls,err := ioutil.ReadFile(file)
  if err != nil {fmt.Println(err)}
  // build up everything we have in clusters and delete found words from gold word list
  clusters := strings.Split(string(binCls), "\n")
  idx :=0
  for ; idx<len(clusters) ; idx++ {
    fields := strings.Split(clusters[idx], "\t")
    words := strings.Split(fields[2], ", ")
    for _,w := range(words) {
      if _,oo := tmp[w] ; oo {
	if _,ok := cluster[w] ; !ok {
	  cluster[w] = make(map[int]int)
	}
	cluster[w][idx]++
	delete(tmp, w)
      }
    }
  }
//   fmt.Println(tmp)
/*  add missing terms back as singletons */
  for term,_ := range(tmp) {
    cluster[term] = map[int]int{idx:1}
    idx++
  }
  return
}






// Return only the most frequent word in each cluster (based on absolute counts)
func (data *NounCluster) ExtractTopWords(numTop, topN int, prnt bool) (outPairs map[int]util.KeyValuePairs) {
  cluster := data.GetClusterAbsoluteFreq(numTop, "soft")
  outPairs = make(map[int]util.KeyValuePairs, len(cluster))
  cntTable := make(map[int]map[string]int)
  for _,cl := range(cluster) {
    cntTable[cl.Name] = make(map[string]int)
    for _, mem := range(cl.Members) {
      cntTable[cl.Name][mem]++
    }
    pairs := util.SortKeysByValues(cntTable[cl.Name])
    if topN <= len(pairs) {
      outPairs[cl.Name] = pairs[len(pairs)-topN:len(pairs)]
    } else {
      outPairs[cl.Name] = pairs
    }
    if prnt {
      fmt.Println("\n\n", cl.Name)
      for _,pair := range(outPairs[cl.Name]) {
	fmt.Print(pair.Key, "(", pair.Value, ")  ")
      }
    }
  }
  return
}



// return the most salient topic for each word (based on absolute counts)
func (data *NounCluster) ExtractTopTopics(topN int) {
  for instance, cl := range(*data) {
    if len(cl) > topN {
      counts := make([]int, len(cl))
      idx :=0 
      for _,count := range(cl) {
	counts[idx] = count
	idx++
      }
      sort.Ints(counts)
      minCount := counts[len(counts)-topN]
      if counts[len(counts)-topN] > 20 {
	for topic,count := range(cl) {
	  if count < minCount {
	    delete((*data)[instance], topic)
	  }
	}
      } else {
	(*data)[instance] = map[int]int{}
      }
    }
  }
}




/* take normal word-indexed representation
 * return ID:[]member representation
 */
func (data NounCluster) GetClusterAbsoluteFreq(numTop int, mode string) (out2 []cluster) {
  tmp := make(map[int][]string)
  for instance,topics := range(data) {
    for id,count := range(topics) {
      if mode=="soft" {
	for i:=0 ; i<count ; i++ {
	  tmp[id] = append(tmp[id], instance)
	}
      } else {
	tmp[id] = append(tmp[id], instance)
      }
    }
  }
  out := make([]cluster, numTop)
  for id, members := range(tmp) {
    out[id].Name = id
    out[id].Members = members
  }
  out2 = make([]cluster, len(out))
  wIdx := 0
  for _, cl := range(out) {
    if len(cl.Members) != 0 {
      out2[wIdx] = cl
      wIdx++
    }
  }
  return out2[:wIdx]
}





// Delete all nouns which are not in gold vocabulary
func (cluster NounCluster) Filter (goldVocab map[string]int) {
  for key, _ := range(cluster) {
    if _, ok := goldVocab[key] ; !ok {
      delete(cluster, key)
    }
  }
}




// print id:member map
func (data NounCluster) PrintCl2Word(numTop int) {
  out := data.GetClusterAbsoluteFreq(numTop, "soft")
  cnt := 0
  for _,cl := range(out) {
    lastW := cl.Members[0]
    cnt = 0
    fmt.Print("\n", cl.Name, "  ", len(cl.Members), "    ")
    for _, mem := range(cl.Members) {
      if mem != lastW {
	fmt.Print(lastW, "(", cnt, ")", "  ")
	lastW = mem
	cnt = 0
      }
      cnt++
    }
    fmt.Print(lastW, "(", cnt, ")", "  ")
  }
  fmt.Println("\n")
}




// print id:member map
func (data NounCluster) String(numTop int) (str string) {
  out := data.GetClusterAbsoluteFreq(numTop, "soft")
  cnt := 0
  for _,cl := range(out) {
    lastW := cl.Members[0]
    cnt = 0
    str += fmt.Sprint("\n", cl.Name, "  ", len(cl.Members), "    ")
    for _, mem := range(cl.Members) {
      if mem != lastW {
	str += fmt.Sprint(lastW, "(", cnt, ")", "  ")
	lastW = mem
	cnt = 0
      }
      cnt++
    }
    str += fmt.Sprint(lastW, "(", cnt, ")", "  ")
  }
  str += "\n"
  return 
}







func (data *NounCluster) PrintSingleCluster(id int) {
  outStr := ""
  outStr += fmt.Sprintf("%d:   ", id)
  fd := false
  for word,tps := range(*data) {
    for tp,_ := range(tps){
      if tp==id {
	outStr += fmt.Sprintf("%s  ", word)
        fd = true
      }
    }
  } 
  if !fd {return}
  fmt.Println(outStr)
}



func (data NounCluster) PrintWordTopicCounts() {
  for instance,topics := range(data) {
    fmt.Println()
    fmt.Print(instance)
    for id,cnt := range(topics) {
      fmt.Print(" ", id, ":", cnt)
    }
  }
}






// Storing a go struct containing a NounCluster instance (map word;topic;count) in a file
func (cluster NounCluster) Store(filename string) {
  b := new(bytes.Buffer)
  g := gob.NewEncoder(b)
  err := g.Encode(cluster)
  if err != nil {
    fmt.Println(err)
  }
  fh, eopen := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
  defer fh.Close()
  if eopen != nil {
    fmt.Println(eopen)
  }
  _,ewrite := fh.Write(b.Bytes())
  if ewrite != nil {
    fmt.Println(ewrite)
  }
  //   fmt.Fprintf(os.Stderr, "%d bytes successfully written to file\n", n)
}




// Loading a go struct containing a a NounCluster instance (map word;topic;count) from file
func LoadNounCluster (filename string) (cluster NounCluster) {
  fh, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  dec := gob.NewDecoder(fh)
  err = dec.Decode(&cluster)
  if err != nil {
    fmt.Println(err)
  }
  return
}
