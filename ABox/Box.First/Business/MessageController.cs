using Box.First.Controllers;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.First.Business
{
    /// <summary>
    /// 信息
    /// </summary>
    public class MessageController : BaseController
    {
        [HttpGet, Route("Message")]
        public JsonResult Message()
        {
            //MyConfigs.GetConfig()
            return new JsonResult("这是一条信息");
        }
    }
}
