from faces import faces

def test_hello():
    assert faces("Hello :)") == "Hello 🙂"

def test_goodbye():
    assert faces("Goodbye :(") == "Goodbye 🙁"

def test_hello_goodbye():
    assert faces("Hello :) Goodbye :(") == "Hello 🙂 Goodbye 🙁"
