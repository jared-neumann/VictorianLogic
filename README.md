# Victorian Logic Project
Jared Neumann

This is a repository for code and files associated with the Victorian Logic Project (VLP). The VLP is intended to computationally model different aspects of the discipline of logic throughout the long nineteenth century. Currently, it is focused on the timespan between 1826 and 1860, corresponding to the publication of Richard Whately's <i>Elements of Logic</i> and William Whewell's <i>On the Philosophy of Discovery</i> respectively.

Aspects of discipline to be studied include, but are not limited to: citation analysis, the distribution and evolution of topics over time, and semantic change in <i>technical</i> terms.

### The Nature of the Corpus

Since the project is in its early stages of development, the corpus is small. However, it is exhaustive (as far as I know) given its parameters. It is comprised only of the first editions of monographs explicitly on logic (save for two exceptions which are taken to be a special case: Whewell (1858) and (1860)). In the future, the type of medium will be expanded to include journal and encyclopedia articles, and the time period will be extended in both directions.

The original, unprocessed contents of the corpus can be found in [corpus_original.zip](https://github.com/janeumanIU/LING-L545/blob/master/Project_Victorian_Logic/corpus_original.zip). It contains 44 text files corresponding to unique monographs, and contains a total of 3,755,518 tokens. They were all taken from [Internet Archive](https://archive.org/) and were digitized by Google. This, together with the nature of 19th-century typeface, means that the texts are pretty messy. A worthwhile task would be to improve on the quality of the original corpus. One potential way to do this without improving the scans, OCR, etc., would be to segment the tokens using a dictionary and replace words within a certain edit distance threshold, or join tokens that would make a full word if they were combined. I did not attempt this because I thought that it might generate as many errors as it corrected. But there is probably a smart way to do it.

### Processing the Corpus

The corpus was processed to remove most non-alphabetic characters and clean up unnecessary whitespace, etc., as described in the included [bash script](https://github.com/janeumanIU/LING-L545/blob/master/Project_Victorian_Logic/misc/clean_text.sh#L4):
```
tr -d '[:punct:]' | tr -d '[:digit:]' | tr -d '[«»@#$%^*}{|£¢§]' | tr '[:upper:]' '[:lower:]' | tr '[\n\r]' ' ' | tr -s ' '
```
Punctuation was removed so that the corpus works better with the intended tools, which do not rely on annotations of any kind.

Then, a [customized list of stopwords](https://github.com/janeumanIU/LING-L545/blob/master/Project_Victorian_Logic/misc/stopwords_mallet_en.txt) (based on the Mallet list for English) were used with the included [Python script](https://github.com/janeumanIU/LING-L545/blob/master/Project_Victorian_Logic/misc/remove_stopwords.py). There is no doubt that this list could be improved, but it is already an improvement over the list included with Python's NLTK. The processed corpus with the cleaned text and removed stopwords can be found in [corpus_processed.zip](https://github.com/janeumanIU/LING-L545/blob/master/Project_Victorian_Logic/corpus_processed.zip).  It contains a total of 1,492,672 tokens. Here is a list of the top 15 tokens in the processed corpus with frequencies (per 10,000 words):

<details><summary>See table</summary>
 
Rank | Word        | Freq.
:---:|:-----------:|:-----:
1    | logic       | 45.00
2    | general     | 38.70
3    | science     | 37.89
4    | nature      | 34.21
5    | form        | 30.98
6    | term        | 30.82
7    | man         | 30.76
8    | subject     | 30.65
9    | knowledge   | 29.97
10   | true        | 29.85
11   | mind        | 29.41
12   | proposition | 27.68
13   | terms       | 27.62
14   | reasoning   | 27.60
15   | laws        | 26.91

</details>

## Modeling Semantic Change in Jargon
This is the first aspect to be explored. I am interested primarily in changes with respect to the semantics of <i>jargon</i> or <i>technical</i> terms on the assumption that conceptual changes in the discipline correlate with changes in its taxonomy. 

There are many options for modeling semantic change out there right now. Several notable examples are outlined in [Tang (2018)](https://arxiv.org/abs/1801.09872). For my purposes, I chose to use the SCAN model explained in [Frermann and Lapata (2016)](https://www.aclweb.org/anthology/Q16-1003.pdf). SCAN is a Bayesian model of diachronic meaning change which the creators experimentally show worked on-par with other systems available at the time of publication. It also offers a number of advantages for use by historians. It models word senses as probability distrubitions over words, and is capable of modeling changes in the number of senses, changes in their corresponding representativeness, as well as <i>subtle</i> changes within individual senses. SCAN models semantic change as 'smooth', based on time-sliced distributions over vocabulary. Unfortunately, the current corpus is small, and not every time-slice is appropriately fleshed out, but that will change with the development of this project and there are still interesting things to learn with SCAN. Additionally, the source code is user-friendly, and can be found in Lea Frermann's repository for it [here](https://github.com/ColiLea/scan).

SCAN requires three inputs from the user: (1) a target word list for words to be modeled, (2) a corpus of concordances for those words, and (3) a parameters file which gives the model the file paths and sets its variables.

### Choosing the Target Words
Any number of target words could be chosen. Since I am interested in disciplinary change, I chose some prominent words that are integral to the Victorian taxonomy of logic, e.g., <i>logic</i>, <i>reasoning</i>, <i>induction</i>, <i>syllogism</i>, and <i>inference</i>. I would have also included <i>deduction</i>, but it was anomalously under-represented in the corpus, meaning that something about the word probably messed with the OCR. Here are the ranks of the terms in the processed corpus:

<details>
 
 <summary>See table</summary>

Rank  | Word      | Freq.
:----:|:---------:|:-----:
1     | logic     | 45.00
14    | reasoning | 27.60
28    | syllogism | 20.90
53    | induction | 16.06
128   | inference | 9.29
...   | ...       | ...
653   | deduction | 2.65

</details>

It is an interesting question whether there is an obvious distinction between these and other terms. One criterion that not all of the above satisfy is that there exists an explicit definition given in the text; but, in the case especially of <i>reasoning</i> and <i>inference</i>, which seem to be jargon terms, this is not always or even typically present. Another potential criterion might be the frequency of the terms relative to a more general, non-disciplinary corpus. I think this would yield a lot of false positives, though, since even non-discipline specific language can be common, or more highly represented, in the discipline without it being strictly <i>jargon</i>, as might be indicated in the word frequencies listed in the previous section.

Another question of interest is whether lemmatizing or stemming the terms (e.g., logic>al|ally, syllogi>sm|ze|stic|stically) gives better or worse results; I'm inclined to think that that <i>jargon</i> terms are better left distinct, as I have left them in this part of the project (for now).

### Creating the Concordances

SCAN requires a concordance file with instances of the target word. The context window is set by the user. To generate concordances, I used Laurence Anthony's [AntConc (Version 3.5.8)](https://www.laurenceanthony.net/software/antconc/) and set the context to 5L and 5R. This generates a set of concordances in the following format (with an example):

 Hit   |        Keyword in Context (KWIC) with 2L/2R                | Origin File
:-----:|:----------------------------------------------------------:|:-------------------------:
...    | ...                                                        | ...
4873   | knowledge science **logic** propositions judgments         | 1850_Field_Analogy.txt
...    | ...                                                        | ...

These original concordances are contained in the [conc_original.zip](https://github.com/janeumanIU/VictorianLogic/blob/master/conc_original.zip) file. However, SCAN requires the following format, where the date of publication comes first, then a tab character, then the concordance:


 Year  | Delimiter |        Keyword in Context (KWIC) with 2L/2R                
:-----:|:---------:|:----------------------------------------------------:
...    | ...       | ...                                                
1850   | \t        | knowledge science **logic** propositions judgments         
...    | ...       | ...                                                

The translation was done with a simple Python script contained in [antconc_to_scan.py](https://github.com/janeumanIU/VictorianLogic/blob/master/misc/antconc_to_scan.py). Furthermore -- another methodological issue -- it may be beneficial to avoid over-representing concordances from texts that are simply longer than others. One way to do this is to put a cap on the number of concordances drawn from any given text. I did not do that for for now, as I thought it might resolve itself given a larger corpus. The new concordances are found in the [input](https://github.com/janeumanIU/VictorianLogic/tree/master/input) folder. 

### Running the SCAN Model

### Interpreting the Output

### Visualizing the Output

## Representing Influence by Citation Analysis

Forthcoming.

## Modeling Topics-over-Time (TOT)

Forthcoming.
