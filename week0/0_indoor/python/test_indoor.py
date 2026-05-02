from indoor import indoor

def test_HELLO():
    output = indoor("HELLO")
    assert output == "hello"

def test_THIS_IS_CS50():
    output = indoor("THIS IS CS50")
    assert output == "this is cs50"

def test_50():
    output = indoor("50")
    assert output == "50"
    assert isinstance(output, str)