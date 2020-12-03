#!/bin/bash

# Vars
REPO_URL="github.com/doliG/welldone"

# Errors
ERR_GO_NOT_FOUND="Ooops, it looks like you don't have Go installed on your machine. Please visit https://golang.org/doc/install to install, then try again."
ERR_ALIAS_FILES_NOT_FOUND="Error: \".bashrc\" and \".zshrc\" not found in \"$HOME\". Please add your alias manually."

# Misc
INSTALL_SUCCESS="🚀 Welldonne sucessfuly installed. Try to run 'welldone add \"My first todo\"' to check."

# 1. Check requirements
if ! command -v go version &> /dev/null
then
	echo $ERR_GO_NOT_FOUND
	exit
fi

# 2. Install package
go get -u $REPO_URL
go install $REPO_URL

# 3. Ask for alias
read -p "Add an alias 't' for 'welldone' ? (y/N) " wantAlias
if [ "$wantAlias" == "y" ]; then
	# Might be converted into array if we want to add more cases
	BASHRC="$HOME/.bashrc"
	ZSHRC="$HOME/.zshrc"

	if [ -f "$BASHRC" ]; then
    	echo "alias t=\"welldone\"" >> "$BASHRC"
		echo "Ok, we've added 'alias t=\"welldone\"' at the end of you \"$BASHRC\""
	elif [ -f "$ZSHRC" ]; then
    	echo "alias t=\"welldone\"" >> "$ZSHRC"
		echo "Ok, we've added 'alias t=\"welldone\"' at the end of you \"$ZSHRC\""
	else
		echo $ERR_ALIAS_FILES_NOT_FOUND
	fi
fi

echo $INSTALL_SUCCESS

# Todo: ask if you want to install an alias "t" instead of "welldone"
# - check if .bashrc || .zshrc
# - ask to the user if he wants to
# - write in file echo "alias t="welldone"" >> .{prefix}rc

# Todo: try / catch error, and ask for people to open issue on gitlab

# Notes: clean installation
# go clean -i github.com/doliG/welldone