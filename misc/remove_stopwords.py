########################################
# Remove stopwords from all text files #
# in directory given a custom stopword #
# list that is also a text file. All   #
# input is obtained via the program.   #
# The program also creates a new       #
# output file with the prefix 'c_'     #
# while preserving the original file.  #
########################################

import os

#Get the stopword list path from user input.
stopwords_file = input("Please give the path to the stopwords list file: ")

#Get the corpus directory from user input.
corpus_dir = input("Please give the directory of the corpus to be processed: ")

#Get the encoding format of the corpus.
enc_format = input("Please specify the encoding format (default with no input is UTF-8): ")

#But set a default for the empty string.
if enc_format == '':
	enc_format = 'utf-8'

#Read the stopword list into a python list without any newline characters.
with open(stopwords_file, 'r') as s:
	stopwords = [line.rstrip() for line in s.readlines()]
s.close()

#Iterate through the files in the given directory.
for file in os.listdir(corpus_dir):

	#First, check to see if the file is a text file.
	if file.endswith(".txt"):
		#Then, open that file using the specified encoding.
		with open(file, 'r', encoding = enc_format) as input_file, open('c_' + file, 'w+') as output_file:
			#And read the lines of that file into a list.
			for line in input_file.readlines():
				#Finally, tokenize the line, convert it back into a space-delimited string without the stopwords,
				#and print the new line to the output file.
				print(" ".join([word for word in line.split() if word not in stopwords]), file = output_file)
		input_file.close()
		output_file.close()
	#Skip files without the text file extension.
	else:
		continue
