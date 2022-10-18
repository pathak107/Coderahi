import { ThemeProvider } from 'next-themes'
import '../styles/globals.css'

function MyApp({ Component, pageProps }) {
  return <ThemeProvider themes={["light", "night"]}>
      <Component {...pageProps} />
  </ThemeProvider>
  
}

export default MyApp
