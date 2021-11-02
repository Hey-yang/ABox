using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.First.Controllers
{
    public class OkController : BaseController
    {
        /// <summary>
        /// 心跳检查方法
        /// </summary>
        /// <returns></returns>
        [HttpGet]
        public ContentResult Get()
        {
            Console.WriteLine($"心跳检查：", DateTime.Now);
            //return Ok();
            return Content("OK");
        }
    }
}
