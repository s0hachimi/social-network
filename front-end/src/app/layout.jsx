
import "./globals.css";




export const metadata = {
  title: "Social Network",
}

export default function RootLayout({ children }) {

  return (
    <html lang="en">
      <body className="">
        {children}
      </body>
    </html>
  );
}
