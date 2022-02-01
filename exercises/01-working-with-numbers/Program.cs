using System;

namespace StudentAverage
{
    class Program
    {
        static void Main(string[] args)
        {
            var mark1 = 75;
            var mark2 = 85;
            var mark3 = 90;
            
            // Casting - Sum of Integers to Float
            var mean = (float)(mark1 + mark2 + mark3) / 3;

            Console.WriteLine("Your average is: {0:0.00}", mean);
            //Console.WriteLine(mean.GetType());
        }
    }
}
