############################################
# Iterates through all text files in a     #
# directory, performs the cleaning on each #
# one, and writes to a new file prefixed   #
# 'c_'.	     				   #
############################################

for f in *.txt
do
	echo "Processing file $f..."
	cat $f | tr -d '[:punct:]' | tr -d '[:digit:]' | tr -d '[«»@#$%^*}{|£¢§]' | tr '[:upper:]' '[:lower:]' | tr '[\n\r]' ' ' | tr -s ' ' > c_$f
done

