""" 
This is the implementation of the exervise put forward by Andrej Karpathy 
Here is the link of the complete exercise https://github.com/karpathy/minbpe/blob/master/exercise.md

Implementing the class of the BasicTokenizer with the following core function:
    -- def train(self, text, vocab_size, verbose=False)
    -- def encode(self, text)
    -- def decode(self, ids)
"""
from typing import List
from utils.getStats import get_stats
from utils.mintTokens import mint_tokens

class basicTokenizer():

    def __init__(self):
        self.startTokenID = 256
        self.merges = {}
        self.vocab = {idx:bytes([idx]) for idx in range(256)}
    
    # Creating a train method for the tokenizer using BPE algorithm
    def train(self, 
            text: str,
            vocab_size: str,
            verbose: bool = False):
        
        """
            Training of the BPE based tokenizer consists of the following steps:
                -- Step # 01: Convert the corpus into bytes that is the utf-8 encoding
                -- Step # 02: Find the stats of the pairs that defines the occurences of the consecutive pairs
                -- Step # 03: Pick the most frequent occurence of the pair
                -- Step # 04: Start minting new tokens and adding them to the merges dictionary containing PAIRS as keys and TOKENIDS as values
                -- Step # 05: Compute the vocabulary with the INDEX being the keys and the BYTES as the values
                -- Step # 06: Store the merges and the vocabularies for further use in encoding and decoding
        """

        # Getting bytes
        tokens = list(text.encode("utf-8"))

        # Train to the point when we reach the 
        while len(self.vocab) < vocab_size:
            # Getting the count statistics and finding the most frequent pair
            if verbose:
                print("*"*100)
                print("*"*100)
                print("Getting stats for the corpus...")
            count_stats = get_stats(tokens)

            if verbose:
                print(f"\nStats : {count_stats}")
            most_frequent_pair = max(count_stats, key=count_stats.get)

            if verbose:
                print(f"\nMost Frequent Pair : {most_frequent_pair} with a count of : {count_stats[most_frequent_pair]}")
                print(f"\nMinting token for tokenID : {self.startTokenID}")

            # Mint tokens based on the frequent pair
            tokens, newMerge = mint_tokens(tokens,
                                            most_frequent_pair,
                                            self.startTokenID)
            
            if verbose:
                print(f"\nNew token minted for tokens {most_frequent_pair}")
                print(f"After minting tokens, we get : {tokens}")    

            # Update lookups
            self.merges = self.merges | newMerge
            self.vocab[self.startTokenID] = self.vocab[most_frequent_pair[0]] + self.vocab[most_frequent_pair[1]]

            # Updating the newTokenID to be appended
            self.startTokenID += 1
        
        if verbose:
                print(f"\n\nFinished training tokenizer")
                print("*"*100)
                print("*"*100)
    

    """ Method to encode the text into tokens """
    def encode(self,
               text: str) -> List[int]:

        tokens = list(text.encode("utf-8"))
        
        while True:
            idx = 0
            newTokens = []
            isMergeLeft = False

            while idx < len(tokens):
                if (idx < len(tokens)-1):
                    pair = tokens[idx], tokens[idx+1]
                    if pair in self.merges:
                        newTokens.append(self.merges[pair])
                        idx += 2

                        # If we have a merge then it can be potentially used to merge with this new token ID
                        isMergeLeft = True
                        continue
                
                newTokens.append(tokens[idx])
                idx += 1

            if not isMergeLeft:
                break
            
            tokens = newTokens[:]
        
        return tokens
    
    """ Method to decode the tokens into text """
    def decode(self,
               tokens: List[int]) -> str:
                
        extractedBytes = [self.vocab[token] for token in tokens]
        byteText = b"".join(byte for byte in extractedBytes)
        text = byteText.decode("utf-8", errors='replace')
        return text