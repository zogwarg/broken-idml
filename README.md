# Windows instructions

1. Download suitable executable from release page https://github.com/zogwarg/broken-idml/releases, and allow the program to run locally.
2. Open powershell console with with Win+R and enter "powershell"
3. Navigate to executable location with cd. eg: "cd ~/Desktop"
4. Run the exe in current dir by entering "./broken-idml-windows-amd64.exe"
5. Select broken idml files
6. Get list of errors in console output
7. Rename your idml file to a zip file eg: "source.idml" -> "source.zip"
8. Double click on the archive file, and navigate to the broken xml files.
9. Drag and drop files for example to desktop.
10. Open the xml files with a text editor, eg: Notepad
11. Manually repair error: eg broken entity sequence " S&D " -> " S&amp;amp;D "
12. Drag all repaired files into the correct location in the archive.
13. Rename the zip file back to being an idml file, eg: "source.zip" -> "source.idml"
14. (Optional) Run "./broken-idml-windows-amd64.exe" again, and there should be empty output for the repaired idml file if there are no errors.
