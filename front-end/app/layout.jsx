import Link from "next/link";
import "./globals.css";



export const metadata = {
  title: "Social Network",
}

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className="">

        <header>
          <nav className="navigation">

            <Link href="/Profile">
              <button>Profile</button>
            </Link>

             <Link href="/">
              <button>Posts</button>
            </Link>

             <Link href="/Groups">
              <button>Groups</button>
            </Link>

             <Link href="/Followers">
             <button>Followers</button>  
            </Link>

             <Link href="/Chats">
               <button>Chats</button>
            </Link>

            <Link href="/Notification">
                <button>Notification</button>
            </Link>

          </nav>
        </header>


        {children}


      </body>
    </html>
  );
}
