import { useState, useEffect } from 'react'
import { useTheme } from 'next-themes'
import {FaMoon, FaSun} from 'react-icons/fa'

const ThemeToggle = () => {
  const [mounted, setMounted] = useState(false)
  const { theme, setTheme } = useTheme()

  // useEffect only runs on the client, so now we can safely show the UI
  useEffect(() => {
    setMounted(true)
  }, [])

  if (!mounted) {
    return null
  }

  const toggleTheme = ()=>{
    if (theme==='light'){
        setTheme('night')
    }else{
        setTheme('light')
    }
  }

  return (
    <div className='flex items-center gap-1'>
        <input type="checkbox" className="toggle" onChange={toggleTheme} checked={theme==='light'? false: true} />
        {theme==='light' ? <FaSun /> : <FaMoon/>}
    </div>
    
  )
}

export default ThemeToggle