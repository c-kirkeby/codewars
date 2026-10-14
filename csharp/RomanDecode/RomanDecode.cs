using System;
using System.Collections.Generic;
using System.Linq;

public class RomanDecode
{
  private static readonly Dictionary<char, int> lookup = new(){
    { 'I', 1 },
    { 'V', 5 },
    { 'X', 10 },
    { 'L', 50 },
    { 'C', 100 },
    { 'D', 500 },
    { 'M', 1000 }
  };

  public static int Solution(string roman)
  {
    var chars = roman.ToCharArray();
    var sum = 0;

    for (var i = chars.Length - 1; i >= 0; i--)
    {
      var current = chars[i];
      sum += lookup[current];
      if (i == 0)
      {
        return sum;
      }

      switch (current)
      {
        case 'D':
        case 'M':
          if (roman[i - 1] == 'C')
          {
            sum -= 100;
            i--;
          }
          break;
        case 'C':
        case 'L':
          if (roman[i - 1] == 'X')
          {
            sum -= 10;
            i--;
          }
          break;
        case 'V':
        case 'X':
          if (roman[i - 1] == 'I')
          {
            sum -= 1;
            i--;
          }
          break;
      }
    }
    return sum;
  }
}
