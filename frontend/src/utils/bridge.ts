import * as Fox from '@wails/go/backend/Fox'
import * as App from '@wails/go/main/App'
export * from '@wails/runtime/runtime'


export const Writefile = async (path: string, content: string) => {
  const { flag, data } = await App.Writefile(path, content)
  if (!flag) {
    throw data
  }
  return data
}

export const Readfile = async (path: string) => {
  const { flag, data } = await App.Readfile(path)
  if (!flag) {
    throw data
  }
  return data
}

export const Notification = async (title: string, message: string, icon = '') => {
  const { flag, data } = await App.Notification(title, message, icon)
  if (!flag) {
    throw data
  }
  return data
}

export const Getcwd = async () => {
  const { flag, data } = await App.Getcwd()
  if (!flag) {
    throw data
  }
  return data
}

export const Exec = async (path: string, args: string) => {
  const { flag, data } = await App.Exec(path, args)
  if (!flag) {
    throw data
  }
  return data
}

export const GetUserCwd = async () => {
  const { flag, data } = await App.GetUserCwd()
  if (!flag) {
    throw data
  }
  return data
}

export const StartScan = async () => {
  const allData = new Array<object>
  const { content } = await Fox.StartScan()

  for (const key in content) {
    if (Object.prototype.hasOwnProperty.call(content, key)) {
      const element = content[key];
      allData.push(...element)
    }
  }

  return allData
}

export const SetNetWork = async (ip: string) => {
  const setInfo = await Fox.SetNetWork(ip)
  return setInfo
}

export const RestoreNetWork = async () => {
  return await Fox.RestoreNetWork()
}

export const CheckNetWork = async () => {
  return await Fox.CheckNetWork()
}

export const GetClientIP = async () => {
  return await Fox.GetClientIP()
}