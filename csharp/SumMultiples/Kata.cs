public static class Kata
{
  public static Func<int, int> GetMultiples = (value) =>
  {
    int sum = 0;
    for (int i = value - 1; i > 0; i--)
    {
      if (i == 0) return sum;
      if (i % 5 == 0 || i % 3 == 0) sum += i;
    }
    return sum;
  };
}
