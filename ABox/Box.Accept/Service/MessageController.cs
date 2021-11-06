using Box.Accept.Controllers;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.Accept.Service
{
    /// <summary>
    /// 信息
    /// </summary>
    public class MessageController : BaseController
    {
        [HttpGet, Route("Message")]
        public JsonResult Message()
        {
            return new JsonResult("接收了一个信息！");
        }
    }
}
