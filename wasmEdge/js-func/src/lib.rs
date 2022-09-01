use wasmedge_quickjs::*;
#[allow(unused_imports)]
use wasmedge_bindgen::*;
use wasmedge_bindgen_macro::*;

#[wasmedge_bindgen]
pub fn run(s: String) -> String {
  let mut ctx = Context::new();

  let code = include_str!("js/run.js");
  let r = ctx.eval_global_str(code);
  if let JsValue::Function(f) = r {
    let param = ctx.new_string(&s);
    let mut argv = vec![param.into()];
    let r = f.call(&mut argv);
    match r {
      JsValue::String(s) => {
        return s.to_string();
      }
      JsValue::Int(i) => {
        return i.to_string();
      }
      JsValue::Float(f) => {
        return f.to_string();
      }
      JsValue::Bool(b) => {
        return b.to_string();
      }
      JsValue::BigNum(b) => {
        return b.to_int64().to_string();
      }
      _ => {
        return String::new();
      }
    }
  }

  return String::new();
}