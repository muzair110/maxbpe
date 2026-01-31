# MaxBPE: A Go Implementation of Byte Pair Encoding

Welcome to **MaxBPE**, a minimal yet robust implementation of the Byte Pair Encoding (BPE) algorithm written in Go. This project is a tribute to the logic used by modern LLMs (like GPT-4) to turn raw text into digestible tokens.



## 🚀 The Core Mission
The goal of this repo is to implement the core BPE logic—training, encoding, and decoding—while navigating the unique way Go handles raw bytes versus UTF-8 strings.

This implementation is inspired by **Andrej Karpathy's `minbpe`**. If you are looking to master the theory behind tokenization, I highly recommend checking out his exercise guide:

👉 **[Andrej Karpathy's minbpe Exercise Link](https://github.com/karpathy/minbpe/blob/master/exercise.md)**

---

## 🛠 Features
* **Byte-Level Training**: Learns merges based on raw byte frequencies.
* **Safe Decoding**: Properly handles multi-byte UTF-8 sequences (like Emojis and Korean characters) by joining raw bytes before interpretation.
* **CLI Ready**: Built-in support for command-line arguments using Go's `flag` package.
* **Pure Go**: Zero external dependencies for the core algorithm.

---

## 📖 The "Lifelong Lesson" (Bytes vs. Runes)
One of the key challenges in porting BPE from Python to Go is how strings are handled. In Go, `string(236)` results in a Unicode character (`ì`), whereas BPE requires the literal raw byte. This library implements a "Byte-Safe" fix:
1.  **Cast to Byte**: `byte(idx)`
2.  **Wrap in Slice**: `[]byte{...}`
3.  **Convert to String**: `string(...)`

This ensures that your data remains raw and uncorrupted during the merge process, allowing multi-byte characters to be reconstructed perfectly during decoding.

---

## 💻 Usage

### Prerequisites
Make sure you have [Go](https://go.dev/doc/install) installed (version 1.18+ recommended).

### Running the Tokenizer
To train and run the tokenizer on a custom string, use the following command:

```bash
go run main.go -verbose -text "hello world!!!? (안녕하세요!) lol123 😉" -vocab 260
