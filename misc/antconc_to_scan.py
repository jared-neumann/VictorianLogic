####################################
# This program iterates through    #
# a directory of concordance files #
# in AntConc output format. It     #
# then extracts information and    #
# reorganizes it into SCAN input   #
# format, writing to a new file.   #
# All input is obtained through    #
# the program. It requires the     #
# original filenames to contain    #
# the prefix "[YEAR]_".            #
####################################

import os

#Get the directory path from the user.
conc_dir = input("Please enter the concordance directory absolute path: ")

#Get the encoding format from the user.
enc_format = input("Please enter the encoding format of the directory files (default with no input it UTF-8): ")

#But set a default for the empty string.
if enc_format == '':
	enc_format = 'utf-8'

#Iterate through the files in the given directory.
for file in os.listdir(conc_dir):

	#First, check to see if the file is a text file.
	if file.endswith(".txt"):
		#Then, open that file using the specified encoding.
		with open(file, 'r', encoding = enc_format) as input_file, open("scan_" + file, 'w+') as output_file:
			#Iterate through the lines of that file.
			for line in input_file:
				#And split them with the TAB delimiter.
				tokens = line.split("\t")
				#Also, split the filename with the UNDERSCORE delimiter.
				tokens[3] = tokens[3].split("_")
				#And print the following pieces to the new file: [YEAR]\t[CONCORDANCE].
				print(tokens[3][0] + "\t" + tokens[1] + tokens[2] + "\n", file = output_file)
		input_file.close()
		output_file.close()
	else:
		continue
