# This function returns the stats about the number of occurences of the pairs in the given text corpus
from typing import Dict
from typing import List

def get_stats(tokens: List[int],
            counts: Dict | None = None) -> Dict:
    
    counts = {} if counts == None else counts

    # Counting the consecutive pairs
    for pair in zip(tokens, tokens[1:]):
        counts[pair] = counts.get(pair,0) + 1
    
    return counts