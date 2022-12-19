export const fmtTitle = (title, now) => {
  const reg = /\$\{(.+?)\}/
  const reg_g = /\$\{(.+?)\}/g
  const result = title.match(reg_g)
  // console.log(title)
  if (result) {
    // console.log(result)
    result.forEach((item) => {
      const key = item.match(reg)[1]
      const value = now.params[key] || now.query[key]
      title = title.replace(item, value)
    })
  }
  return title
}
