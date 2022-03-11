using Box.Apollo.Controllers;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.Apollo.Business
{
    
    public class TestController : BaseController
    {
        private readonly IConfiguration _configuration;

        public TestController(IConfiguration configuration)
        {
            _configuration = configuration;
        }
        [HttpGet, Route("Message")]
        public JsonResult Message()
        {
            var sss = _configuration["Text"];
            return new JsonResult("这是一条信息" + sss);
        }
    }
}
