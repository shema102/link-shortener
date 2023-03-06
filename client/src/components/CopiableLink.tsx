import React, {FC, useCallback} from "react";
import './CopiableLink.css'

const CopiableLink: FC<{ url: string }> = ({url}) => {
  const ref = React.useRef<HTMLHeadingElement>(null)

  const onClick = useCallback(() => {
      ref.current?.classList.add('link-animated')

      setTimeout(() => {
        ref.current?.classList.remove('link-animated')
      }, 1000)

      navigator.clipboard.writeText(url)
    }, [url]
  )

  return (
    <div>
      <h2 className={"link"} ref={ref} onClick={onClick}>
        {url}
        <div className={"press-to-copy"}>(Press to copy)</div>
      </h2>
    </div>
  )
}

export default CopiableLink
