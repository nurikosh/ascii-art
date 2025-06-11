# ASCII Art Generator

## 📄 Overview

This program converts input text into stylized ASCII art using predefined character templates from a banner file. It supports multi-line input and handles special characters gracefully.

---

## ✨ Features

- Loads custom ASCII character templates from banner files
- Processes multi-line input with `\n` separators
- Generates properly aligned ASCII art output

---

## 🛠️ Installation

1. Ensure you have Go installed (**version 1.16+ recommended**)
2. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/ascii-art-generator.git
Navigate to the project directory:
```bash
cd ascii-art-generator
```
🚀 Usage
Run the program using:
   ```bash
   go run ./cmd [input_string]
   ```
🔹 Arguments
input_string: The text to convert to ASCII art (optional, can also be provided at runtime).

🗂️ Banner Files
The program looks for banner files in the banners/ directory.

The default banner file is standard.txt.

Each character in the banner file is represented by 8 lines of ASCII art.

⚙️ How It Works
1. Banner Loading
Reads character templates from the banner file.

Stores characters in a map for fast access.

2. Text Processing
Splits input text by \n for multi-line support.

Handles empty lines and special characters.

3. ASCII Art Generation
For each input line:

Constructs the 8-line ASCII art row by row.

Horizontally combines character templates.

Prints the completed ASCII art output.

📌 Example
Input:
```bash
go run ./cmd "Hello\nThere"
```
Output:
<pre>
 _    _          _   _
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
</pre>
🔮 Future Enhancements
✅ Support for colored output

✅ Additional banner styles (e.g., shadow, thinkertoy)

✅ Command-line flags for output customization

✅ Interactive input mode

✅ Better error handling for malformed banner files

📦 Dependencies
Uses only the Go standard library:

bufio

fmt

os

strings
