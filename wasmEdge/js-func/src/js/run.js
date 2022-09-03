// (s) => {
// 	return s;
// }

(jsonStr) => {
	let obj= JSON.parse(jsonStr)
	const {deviceId  :code } =obj
	// TODO: Edit the following lines to modify the threshold of temperature and humidity
	let temperatureThreshold = 0;
	let humidityThreshold = 0;
	return JSON.stringify(obj.entity.map(it=>{
	  if(it.eUnit==='%RH')
	  return {
		  code,
		  datetime: it.datetime,
		  name:it.eName,
		  val:it.eValue,
		  unit:'%RH',
		  exception:'湿度' + (it.eValue> temperatureThreshold ? '过高':'正常')
	  }
	  else
	  return {
		  code,
		  datetime: it.datetime,
		  name:it.eName,
		  val:it.eValue,
		  unit:'℃',
		  exception:'温度' +( it.eValue> humidityThreshold ? '过高':'正常')
	  }
	}) )
  }