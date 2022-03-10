using Box.Test.Controllers.Base;
using Consul;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;

namespace Box.Test.Controllers.Business
{
    /// <summary>
    /// 测试
    /// </summary>
    public class TestController : BaseController
    {
        private readonly IHttpClientFactory _httpClientFactory;
        public TestController(IHttpClientFactory httpClientFactory)
        {
            _httpClientFactory = httpClientFactory;
        }
        /// <summary>
        ///  测试
        /// </summary>
        /// <returns></returns>w
        [HttpGet]
        public async Task<IActionResult>  Get()
        {
            using (var consulClient = new ConsulClient(a => { a.Address = new Uri("http://192.168.232.128:8500"); a.Datacenter = "dc1"; }))
            {
                var services = consulClient.Catalog.Service("YangService").Result.Response;
                if (services != null && services.Any())
                {
                    // 模拟随机一台进行请求，这里只是测试，可以选择合适的负载均衡框架
                    var service = services.ElementAt(new Random().Next(services.Count()));

                    var client = _httpClientFactory.CreateClient();
                    var response = await client.GetAsync($"http://{service.ServiceAddress}:{service.ServicePort}/api/Message/Message");
                    var result = await response.Content.ReadAsStringAsync() + service.ServicePort;
                    return new JsonResult(result);
                }
            }
            return new JsonResult("123");
        }
    }
}
