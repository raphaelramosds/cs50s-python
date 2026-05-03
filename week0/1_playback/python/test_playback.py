from playback import playback

def test_This_is_CS50():
    assert playback("This is CS50") == "This...is...CS50"

def test_This_is_our_week_on_functions():
    assert playback("This is our week on functions") == "This...is...our...week...on...functions"

def test_Let_s_implement_a_function_called_hello():
    assert playback("Let\'s implement a function called hello") == "Let\'s...implement...a...function...called...hello"