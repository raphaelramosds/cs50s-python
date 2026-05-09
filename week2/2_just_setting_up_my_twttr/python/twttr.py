import re

def twttr(input: str) -> str:
    input = input.strip()
    # Replace all vowels with an empty string with a regular expression
    return re.sub(r'[aeiouAEIOU]', '', input)