using System.Linq;

public static class Kata
{
  public static bool IsPangram(string str)
  {
    return str.Where(char.IsLetter).Select(char.ToLower).Distinct().Count() == 26;
  }
}
