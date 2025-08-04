
# 📦 Packet Processor CLI

`packet_processor.exe` is a lightweight command-line tool for generating and/or processing sample data from input files.
---

## 📦 Download & Setup

1. Download `packet_processor.exe` from here: [Download Executable](https://github.com/yourhonor1996/optimized_packet_processor/blob/main/packet_processor/packet_proecssor.exe)
2. Place it in a directory of your choice.
3. Open **Command Prompt** (`cmd`) or **PowerShell**.
4. Navigate to the directory containing `packet_processor.exe`.

```sh
cd C:\path\to\packet_processor
---

## 🚀 Usage (Windows)

```cmd
packet_processor.exe [-i input_file] [-o output_file] [-n num_samples]
```

All parameters are optional. If you don't provide them, the application will generate default values.

---

## 🧾 Command-Line Options

| Option | Description                                                                                                                                                                                               |
|--------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `-i`   | (Optional) Path to input file. If not provided, the app generates one.                                                                                                                                    |
| `-o`   | (Optional) Path to output file. If not provided, the app generates one.                                                                                                                                   |
| `-n`   | (Optional) Number of samples to generate or process. Default behavior:<br>• If `-i` is **not** provided → generate **`n`** samples.<br>• If `-i` **is** provided → process **`n`** samples from the file. |
| `-h`   | Help for commands                                                                                                                                                                                         |


---

## 🧪 Examples

### ✅ 1. **Generate and process 100 samples**

```cmd
packet_processor.exe -n 100
```

> This will generate a new input file with 100 samples and write results to an auto-generated output file.

---

### ✅ 2. **Process 50 samples from an existing input file**

```cmd
packet_processor.exe -i data\input.csv -n 50
```

> Reads from `data\input.csv` and processes 50 samples.

---

### ✅ 3. **Process from file and write to custom output**

```cmd
packet_processor.exe -i input.csv -o results.csv -n 200
```

> Reads 200 samples from `input.csv` and writes processed results to `results.csv`.

---

### ✅ 4. **Help command**

```cmd
packet_processor.exe -h
```

> Shows usage help for all available flags.

---

## 📁 Output

* Output files (generated or specified) will contain processed data.

