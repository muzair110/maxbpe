# This function is responsible for minting new tokens based on the most frequent token pair criterion for BPE
from typing import List
from typing import Tuple
from typing import Union
from typing import Dict

def mint_tokens(tokens: List[int],
                most_frequent_pair: Tuple,
                newTokenID: int) -> Union[List[int], Dict]:
    
    """
    Input Arguments:
        tokens: These are the raw tokens of the corpus
        most_frequent_pair: This is the most frequent pair to be searched and minted for a new token
        newTokenID: This is the token ID that should be the first new token ID to start minting the tokens. This parameter is to facilitate
                      the retraining of the tokenizer

    Returns:
        -- First if the List[int] that is the updated tokens with the most_frequent_pair being merged and a new token has been minted
        -- Dict is the merged dictionaries that will be needed            
    """
    idx = 0
    newTokens = []

    while idx < len(tokens):

        # Mint new tokens by tracing the most frequent pair in the token sequence
        if tokens[idx] == most_frequent_pair[0] and (idx < len(tokens) - 1) and tokens[idx + 1] == most_frequent_pair[1]:
            newTokens.append(newTokenID)
            idx += 2
        else:
            newTokens.append(tokens[idx])
            idx += 1
    
    return newTokens, {most_frequent_pair : newTokenID}