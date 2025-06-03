import Link from "next/link";

export default function NavBar() {
    return (
          <header>
          <nav className="navigation">

            <Link href="/Profile">
              <button>Profile</button>
            </Link>

             <Link href="/Posts">
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
    )
}