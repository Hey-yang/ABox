using Box.First.Models;
using Consul;
using Microsoft.AspNetCore.Builder;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.Extensions.Hosting;

namespace Box.First.Helpers
{
    public static class ConsulHelpers
    {
        public static IApplicationBuilder RegisterConsul(this IApplicationBuilder app, IHostApplicationLifetime lifetime, ConsulServiceInfo consulServiceOptions)
        {
            var consulClient = new ConsulClient(x =>
            {
                // consul服务地址，默认安装consul后端口为8500
                x.Address = new Uri(consulServiceOptions.ConsulAddress);
                x.Datacenter = consulServiceOptions.ConsulCenter;
            });
            var registration = new AgentServiceRegistration()
            {
                ID = "yang" + Guid.NewGuid().ToString(),
                Name = consulServiceOptions.ServiceName,// 服务名
                Address = consulServiceOptions.ServiceIP, // 服务绑定IP(也就是你这个项目运行的ip地址)
                Port = consulServiceOptions.ServicePort, // 服务绑定端口(也就是你这个项目运行的端口)
               // Tags = new string[] { }, //标签，用来标记，相关备注
                Check = new AgentServiceCheck()
                {
                    DeregisterCriticalServiceAfter = TimeSpan.FromSeconds(5),//服务启动多久后注册
                    Interval = TimeSpan.FromSeconds(10),//健康检查时间间隔
                    HTTP = consulServiceOptions.ServiceHealthCheck,//健康检查地址
                    Timeout = TimeSpan.FromSeconds(5)
                }
            };
            // 服务注册
            consulClient.Agent.ServiceRegister(registration).Wait();
            // 应用程序终止时，服务取消注册
            lifetime.ApplicationStopping.Register(() =>
            {
                consulClient.Agent.ServiceDeregister(registration.ID).Wait();
            });
            return app;
        }
    }
}
