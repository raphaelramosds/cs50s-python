from einstein import einstein

def test_1():
    assert einstein(1) == 90_000_000_000_000_000

def test_14():
    assert einstein(14) == 1_260_000_000_000_000_000

def test_50():
    assert einstein(50) == 4_500_000_000_000_000_000
