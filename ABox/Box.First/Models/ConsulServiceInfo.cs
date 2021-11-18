using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Box.First.Models
{
    /// <summary>
    /// Consul配置模型类
    /// </summary>
    public class ConsulServiceInfo
    {
        /// <summary>
        /// 服务名称
        /// </summary>
        public string ServiceName { get; set; }
        /// <summary>
        /// 服务IP
        /// </summary>
        public string ServiceIP { get; set; }
        /// <summary>
        /// 服务端口
        /// </summary>
        public int ServicePort { get; set; }
        /// <summary>
        /// 服务健康检查地址
        /// </summary>
        public string ServiceHealthCheck { get; set; }
        /// <summary>
        /// Consul 地址
        /// </summary>
        public string ConsulAddress { get; set; }
        /// <summary>
        /// Consul 数据中心
        /// </summary>
        public string ConsulCenter { get; set; }
    }
}
