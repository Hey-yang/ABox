using Box.First.Helpers;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.First
{
    public class Program
    {
        public static void Main(string[] args)
        {
            ReleaseCreateHostBuilder(args).Build().Run();
        }
        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.UseStartup<Startup>();
                });
        public static IHostBuilder ReleaseCreateHostBuilder(string[] args)
        {
            var config = new ConfigurationBuilder().AddCommandLine(args).Build();
            string port = config["port"];
            string ip = config["ip"];
            return Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    //webBuilder.UseStartup<Startup>().UseUrls($"http://{HostHelpers.GetHostIp()}:{port}");
                    webBuilder.UseStartup<Startup>().UseUrls($"http://{ip}:{port}");
                });
        }
          
    }
}
