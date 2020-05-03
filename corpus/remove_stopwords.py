import nltk
import string
from nltk.corpus import stopwords

import os
 
for file in os.listdir('/home/antspiderbee/VictorianLogic/corpus/'):
	if file.endswith(".txt"):
		with open(file,'r', encoding='latin-1') as inFile, open('sw_' + file,'w+') as outFile:
			for line in inFile.readlines():
		    		print(" ".join([word for word in line.split() if word not in stopwords.words('english')]), file=outFile)
	else:
		continue
