#while true; do
#  read -p "Would you like to install Smart Calculator (c) acristin? " yn
#  case $yn in
#      [Yy]* ) break;;
#      [Nn]* ) exit;;
#      * ) echo "Please answer yes or no.";;
#  esac
#done

YELLOW="\033[1;33m"
VIOLET="\033[0;35m"
RED="\033[0;31m"
END="\033[0m"

echo "$YELLOW Smart Calculator (c) acristin installation$END"

while true; do
  read -p " Please enter the installation path: " instPath
  case $instPath in
      exit ) exit;;
      "") echo "$VIOLET Empty path$END";;
      * ) break;;
  esac
done

while true; do
  read -p " Would you like to create a shortcut? " ifCreateShortcutAnswer
  case $ifCreateShortcutAnswer in
      [Yy]* ) ifCreateShortcut=true; break;;
      [Nn]* ) ifCreateShortcut=false; break;;
      * ) echo "$VIOLET Please answer yes or no.$END";;
  esac
done

if [ $ifCreateShortcut == true ]; then
    while true; do
      read -p " Please enter the shortcut path: " shortcutPath
      case $shortcutPath in
          "") echo "$VIOLET Empty path$END";;
          * ) break;;
      esac
    done
fi

if [ ! -d "$instPath" ]; then
  echo "$RED Path $instPath does not exist. Installation stopped.$END"
  exit
fi

if [ $ifCreateShortcut == true ]; then
  if [ ! -d "$shortcutPath" ]; then
    echo "$VIOLET Path $shortcutPath does not exist. Shortcut will not be created.$END"
    ifCreateShortcut=false
    shortcutPath=""
  fi
fi

newInstPath="$instPath/SmartCalc"
mkdir -p "$newInstPath"
cp -r assets cmd internal scripts go.mod go.sum Makefile README.md "$newInstPath"
#pwdVar=$PWD
cd "$newInstPath/cmd"
#ls */** | grep Icon
fyne package -icon ../assets/Icon.png -name "SmartCalc"
#pwd
#echo "here path" "$instPath" "shortcut" "$ifCreateShortcut" "path" "$shortcutPath"
