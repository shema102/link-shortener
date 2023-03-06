import {useCallback, useState} from "react";

const getShortUrl = async (url: string, userId: string = "anonymous"): Promise<string> => {
  const response = await fetch('/api/shorten', {
    method: 'POST',
    body: JSON.stringify({
      url,
      userId
    })
  })

  const data = await response.json()

  return data.shortUrl
}

const addHttps = (url: string): string => {
  if (!/^(f|ht)tps?:\/\//i.test(url)) {
    url = "https://" + url
  }
  return url
}

const useShortUrl = () => {
  const [shortUrl, setShortUrl] = useState<string | null>(null)

  const fetchShortUrl = useCallback(async (url: string) => {
    setShortUrl(null)

    const shortUrl = await getShortUrl(addHttps(url))

    setShortUrl(shortUrl)
  }, [])

  return {fetchShortUrl, shortUrl}
}

export default useShortUrl
