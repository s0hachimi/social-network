
import NavBar from "../../components/navBar";
import "../globals.css";




export const metadata = {
  title: "Social Network",
}

export default function RootLayout({ children }) {

  return (
    <>
      <NavBar />
      {children}
    </>


  );
}
