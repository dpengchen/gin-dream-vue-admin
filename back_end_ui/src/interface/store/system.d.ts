export class SystemInfo {
  public them: string
  //主题颜色
  public colorPrimary: string
  //当前语言
  public lang: string
  public layout: string
  public sidePosition: string
  public showSide: boolean

  constructor() {
    this.them = 'light'
    this.lang = 'zh-CN'
    this.layout = '1'
    this.sidePosition = 'left'
    this.showSide = true
  }
}
