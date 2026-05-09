from twttr import twttr

def test_twitter():
    assert twttr("twitter") == "twttr"

def test_what_s_your_name():
    assert twttr("what's your name?") == "wht's yr nm?"

def test_cs50s():
    assert twttr("CS50's") == "CS50's"