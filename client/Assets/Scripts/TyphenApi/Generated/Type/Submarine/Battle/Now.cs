// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine.Battle
{
    [MessagePack.MessagePackObject]
    [Newtonsoft.Json.JsonObject(Newtonsoft.Json.MemberSerialization.OptIn)]
    public partial class Now : TyphenApi.TypeBase<Now>, TyphenApi.IRealTimeMessage
    {
        [TyphenApi.QueryStringProperty("time", false)]
        [MessagePack.Key("time")]
        [Newtonsoft.Json.JsonProperty("time")]
        [Newtonsoft.Json.JsonRequired]
        public long Time { get; set; }
    }
}
