# This is a simple test script for the tokenizer class
import argparse
from src.basicTokenizer import basicTokenizer

def main(text: str, vocab_size: int, verbose: bool):
    tokenizer = basicTokenizer()
    tokenizer.train(text,
                    vocab_size,
                    verbose)
    
    # Trying the encoder and the decoder's output
    print("\n\nTesting the Decoded(Encoded) string with the original string. Are these the same...?")
    # print(tokenizer.decode(tokenizer.encode(text)) == text)
    print(f'\nText : {text}')
    print(f'Encoded string : {tokenizer.encode(text)}')
    print(f'Decoded string : {tokenizer.decode(tokenizer.encode(text))}')
    print(f'Are they equal : {tokenizer.decode(tokenizer.encode(text)) == text}')


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='For the input text corpus whether it be a file or text in the terminal')
    parser.add_argument('--text','-t', help='Input text in case file is not present')
    parser.add_argument('--vocab_size', type=int, help='Size of the vocabulary after training')
    parser.add_argument('--verbose','-v', help='Verbose', action='store_true')
    args = parser.parse_args()

    main(args.text, args.vocab_size, args.verbose)