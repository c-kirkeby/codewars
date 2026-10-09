using System;
using System.Linq;
using System.Collections.Generic;

public static class SumSquaredDivisors
{
  public static string ListSquared(long m, long n)
  {
    var pairs = new List<(long, long)>();

    for (long i = m; i <= n; i++)
    {
      var sum = GetDivisors(i).Sum(x => x * x);

      if (IsSquare(sum))
      {
        pairs.Add(new(i, sum));
      }
    }
    var values = pairs.Select(pair => $"[{pair.Item1}, {pair.Item2}]");
    return $"[{string.Join(", ", values)}]";
  }

  static List<long> GetDivisors(long n)
  {
    List<long> divisors = [];

    for (int i = 1; i <= Math.Sqrt(n); i++)
    {
      if (n % i == 0)
      {
        divisors.Add(i);
        if (i != n / i)
        {
          divisors.Add(n / i);
        }
      }
    }
    divisors.Sort();
    return divisors;
  }

  static bool IsSquare(long x)
  {
    var sqrt = (long)Math.Sqrt(x);
    return sqrt * sqrt == x;
  }
}
