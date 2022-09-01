// (s) => {
// 	return s;
// }

(jsonStr) => {
	let obj= JSON.parse(jsonStr)
	const {deviceId  :code } =obj
	return JSON.stringify(obj.entity.map(it=>{
	  if(it.eUnit==='%RH')
	  return {
		  code,
		  name:it.eName,
		  val:it.eValue,
		  unit:'%RH',
		  exception:'湿度' + (it.eValue>85? '过高':'正常')
	  }
	  else
	  return {
		  code,
		  name:it.eName,
		  val:it.eValue,
		  unit:'℃',
		  exception:'温度' +( it.eValue>28? '过高':'正常')
	  }
	}) )
  }