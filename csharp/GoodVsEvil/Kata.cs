using System;
using System.Collections.Generic;
using System.Linq;

public static class Kata
{
  static readonly Dictionary<string, int[]> worth = new()
  {
    { "good", new int[]{1, 2, 3, 3, 4, 10} },
    { "evil", new int[]{1, 2, 2, 2, 3, 5, 10} }
  };

  public static Func<string, string, string> goodVsEvil = (good, evil) =>
  {
    var result = GetStrength(worth["good"], good) - GetStrength(worth["evil"], evil);

    return result switch
    {
      < 0 => "Battle Result: Evil eradicates all trace of Good",
      > 0 => "Battle Result: Good triumphs over Evil",
      _ => "Battle Result: No victor on this battle field",
    };
  };
  static int GetStrength(int[] worth, string units) => units.Split(" ").Select((unit, index) => Convert.ToInt32(unit) * worth[index]).Sum();
};
