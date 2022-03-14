using System;
using System.Threading;

namespace Box.Polly
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");

            var policy = PollyHTool.CreatePolly();

            for (int i = 0; i < 100; i++)
            {
                Console.WriteLine($"-------------第{i}次请求-------------");
                policy.Execute(() =>
                {
                    // 从10次开始，正常请求成功
                    if (i < 10)
                    {
                        Thread.Sleep(3000);
                    }
                    else
                    {
                        Console.WriteLine($"{DateTime.Now}：请求成功");
                    }
                });
                Thread.Sleep(1000);
            }
        }
    }
}
