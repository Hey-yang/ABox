using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.Accept.Controllers
{
    /// <summary>
    /// 基类控制器
    /// </summary>
    [ApiController, Route("api/[controller]")]
    public class BaseController : ControllerBase
    {
    }
}
