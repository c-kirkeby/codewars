public static class Kata
{
  public static Func<IEnumerable<int>, int> sumTwoSmallestNumbers = (values) =>
  {
    return values.OrderBy(value => value).Take(2).Sum();
  };
}
