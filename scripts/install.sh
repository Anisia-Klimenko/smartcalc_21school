# Set colors
YELLOW="\033[1;33m"
VIOLET="\033[0;35m"
RED="\033[0;31m"
END="\033[0m"

# Title
echo "$YELLOW Smart Calculator (c) acristin installation$END"

# Ask for installation path
while true; do
  read -p " Please enter the installation path: " instPath
  case $instPath in
      exit ) exit;;
      "") echo "$VIOLET Empty path.$END Press exit to exit installation";;
      * ) break;;
  esac
done

# Ask if shortcut need to be created
while true; do
  read -p " Would you like to create a shortcut? " ifCreateShortcutAnswer
  case $ifCreateShortcutAnswer in
      [Yy]* ) ifCreateShortcut=true; break;;
      [Nn]* ) ifCreateShortcut=false; break;;
      * ) echo "$VIOLET Please answer yes or no.$END";;
  esac
done

# Ask for shortcut path if shortcut need to be created
if [ $ifCreateShortcut == true ]; then
    while true; do
      read -p " Please enter the shortcut path: " shortcutPath
      case $shortcutPath in
          "") echo "$VIOLET Empty path$END";;
          * ) break;;
      esac
    done
fi

# Check if installation path is correct
if [ ! -d "$instPath" ]; then
  echo "$RED Path $instPath does not exist. Installation stopped.$END"
  exit
fi

# Check if shortcut path is correct
if [ $ifCreateShortcut == true ]; then
  if [ ! -d "$shortcutPath" ]; then
    echo "$VIOLET Path $shortcutPath does not exist. Shortcut will not be created.$END"
    ifCreateShortcut=false
    shortcutPath=""
  fi
fi

# Add package name to instalation path and create directory
newInstPath="$instPath/SmartCalc"
mkdir -p "$newInstPath"

# Copy all files to the new directory and change current directory
cp -r assets cmd internal go.mod go.sum Makefile README.md "$newInstPath"
cd "$newInstPath/cmd"

# Run fyne package launcher
fyne package -icon ../assets/Icon.png -name "SmartCalc"

# Move app to the package root
# shellcheck disable=SC2103
cd ..
cp -r cmd/SmartCalc.app .
rm -rf cmd/SmartCalc.app

# Create shortcut
if [ $ifCreateShortcut == true ]; then
  ln -s "$instPath"/SmartCalc/SmartCalc.app "$shortcutPath"
fi
