using Box.AdtExtend.Apollo;
using Com.Ctrip.Framework.Apollo;
using Com.Ctrip.Framework.Apollo.Enums;
using Com.Ctrip.Framework.Apollo.Logging;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.Apollo
{
    public class Program
    {
        public static void Main(string[] args)
        {
            try
            {
                Console.WriteLine("��ʼ����API...");
                var host = CreateHostBuilder(args).Build();
                Console.WriteLine("API�������!");
                host.Run();
            }
            catch (Exception exception)
            {
                Console.ForegroundColor = ConsoleColor.Red;
                Console.WriteLine("API����ʧ��!");
                Console.WriteLine(exception.ToString());
                Console.ResetColor();
            }
        }

        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
            .ConfigureApollo(ConfigFileFormat.Json, "1.Test")
            .ConfigureWebHostDefaults(webBuilder =>
            {
                webBuilder.UseStartup<Startup>();
            });
    }
}
