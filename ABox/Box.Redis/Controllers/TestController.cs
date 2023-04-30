using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace Box.Redis.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class TestController : ControllerBase
    {
        [HttpPost, Route("SetValue")]
        public JsonResult SetValue()
        {
            new RedisHelper().SetValue("test", "1234512");
            return new JsonResult("写入成功test:1234512");
        }

        [HttpGet, Route("GetValue")]
        public JsonResult GetValue()
        {
            var ss =  new RedisHelper().GetValue("test");
            return new JsonResult(ss);
        }
        
        [HttpGet, Route("GetTest")]
        public JsonResult GetTest()
        {
            var ss =  new RedisHelper().GetValue("test");
            return new JsonResult(ss);
        }
    }
}
