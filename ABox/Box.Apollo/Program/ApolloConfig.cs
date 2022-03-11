using Com.Ctrip.Framework.Apollo;
using Com.Ctrip.Framework.Apollo.Enums;
using Com.Ctrip.Framework.Apollo.Logging;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.Apollo.Program
{
    public static class ApolloConfig
    {
        public static IHostBuilder ConfigureApollo(this IHostBuilder builder)
        {
            return builder.ConfigureAppConfiguration((hostingContext, configurationBuilder) =>
            {
                //阿波罗的日志级别调整为最低
                LogManager.UseConsoleLogging(Com.Ctrip.Framework.Apollo.Logging.LogLevel.Trace);
                configurationBuilder.AddApollo(configurationBuilder.Build().GetSection("Apollo"))
                .AddDefault(ConfigFileFormat.Json)
                .AddNamespace("1.Test"); //Apollo中NameSpace的名称,可不使用;
                //configurationBuilder
                //    .AddApollo(hostingContext.Configuration.GetSection("apollo"))
                //    .AddDefault(ConfigFileFormat.Json);
            });
        }
    }
}
