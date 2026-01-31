# MaxBPE: A Python & Go Implementation of the GPT-4 Tokenizer from scratch

Welcome to **MaxBPE**, a minimal yet robust implementation of the Byte Pair Encoding (BPE) algorithm written in Go. This project is a tribute to the logic used by modern LLMs (like GPT-4) to turn raw text into digestible tokens.

## 🚀 The Core Mission

The goal of this repo is to implement the core BPE logic—training, encoding, and decoding—while navigating the unique way Go handles raw bytes versus UTF-8 strings.

This implementation is inspired by **Andrej Karpathy's `minbpe`**. If you are looking to master the theory behind tokenization, I highly recommend checking out his exercise guide:

👉 **[Andrej Karpathy's minbpe Exercise Link](https://github.com/karpathy/minbpe/blob/master/exercise.md)**

---

## 🛠 Features

- **Byte-Level Training**: Learns merges based on raw byte frequencies.
- **Safe Decoding**: Properly handles multi-byte UTF-8 sequences (like Emojis and Korean characters) by joining raw bytes before interpretation.
- **CLI Ready**: Built-in support for command-line arguments using Go's `flag` package.
- **Cross-Language Parity**: Designed to mirror the logic found in Python BPE implementations.

---

## 💻 Usage (Golang)

### Prerequisites

Make sure you have [Go](https://go.dev/doc/install) installed (version 1.18+ recommended).

### Running the Tokenizer

To train and run the tokenizer on a custom string with detailed logging, use the following command:
```bash
go run main.go -verbose -text "hello world!!!? (안녕하세요!) lol123 😉"
```

### Available Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-text` | `"Hello World"` | The input string to train/test on (use quotes for spaces). |
| `-vocab` | `512` | The target vocabulary size (base 256 + merges). |
| `-start` | `256` | The token ID to start assigning to new merges. |
| `-verbose` | `false` | Enable detailed logging of the merge process. |

---

## 🐍 Usage (Python Reference)

If you are using the original Python scripts for comparison (e.g., from minbpe), use the following commands:

**Basic Tokenizer (Core Logic):**
```bash
python -m test.testBasicTokenizer --text "hello world!!!? (안녕하세요!) lol123 😉" --vocab_size 260 --verbose
```

**Regex Tokenizer (GPT-4 Pattern):**
```bash
python -m test.testRegexTokenizer --text "hello world!!!? (안녕하세요!) lol123 😉" --vocab_size 260 --verbose
```

---

## 📖 The "Lifelong Lesson" (Bytes vs. Runes)

One of the key challenges in porting BPE from Python to Go is how strings are handled. In Go, `string(236)` results in a Unicode character (`ì`), whereas BPE requires the literal raw byte. This library implements a "Byte-Safe" fix:

1. **Cast to Byte**: `byte(idx)`
2. **Wrap in Slice**: `[]byte{...}`
3. **Convert to String**: `string(...)`

This ensures that your data remains raw and uncorrupted during the merge process, allowing multi-byte characters (like emojis) to be reconstructed perfectly during decoding.

---

## 🔬 Comparison: Go vs. Python

In the original Python implementation, a complex Regex pattern is used to prevent "category bleeding" (e.g., merging a letter with a punctuation mark).

| Feature | Python minbpe | Go maxbpe |
|---------|---------------|-----------|
| **Logic** | Core BPE | Core BPE |
| **Regex** | GPT-4 Pattern | Core Logic Only* |
| **Data Type** | `bytes` objects | `string` (byte-backed) |
| **Overhead** | Higher (Dynamic) | Lower (Compiled) |

*Note: Go's standard library `regexp` uses RE2, which does not support the possessive quantifiers or lookaheads used in GPT-4's pattern. This repo focuses on the pure implementation of the merging algorithm.*

---

## 🤝 Contributing

Feel free to fork this repo and submit PRs. If you're interested in adding `regexp2` support to fully match the GPT-4 split pattern in Go, that would be a great place to start!

---

**Happy Tokenizing! 🚀**
