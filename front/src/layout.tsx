import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";

// import background from "@/assets/back.jpg";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <SidebarProvider
    //   style={{ backgroundImage: `url(${background})`, width: "100%" }}
    >
      <AppSidebar
      // style={{ background: null }}
      />
      <main>
        <SidebarTrigger />
        <div> {children}</div>
      </main>
    </SidebarProvider>
  );
}
