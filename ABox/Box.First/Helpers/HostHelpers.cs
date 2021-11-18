using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.NetworkInformation;
using System.Net.Sockets;
using System.Threading.Tasks;

namespace Box.First.Helpers
{
    /// <summary>
    /// 本机辅助类
    /// </summary>
    public class HostHelpers
    {
        /// <summary>
        /// 获取宿主机IP
        /// </summary>
        /// <returns></returns>
        public static string GetHostIp()
        {
            //获取本机IP地址(Linux上也可行)
            return NetworkInterface.GetAllNetworkInterfaces()
                           .Select(p => p.GetIPProperties())
                           .SelectMany(p => p.UnicastAddresses)
                           .FirstOrDefault(p => p.Address.AddressFamily == AddressFamily.InterNetwork && !IPAddress.IsLoopback(p.Address))?.Address.ToString();
        }
    }
}
